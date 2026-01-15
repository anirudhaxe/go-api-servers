-- users table (id, name, encrpted_password, role: (admin/user), created at)
--
-- admin - manages agents and compute nodes
-- user - can get list of agents and compute nodes and then he can create a session from the two session id is returned. this session can be updated
--
-- AI agent (id, model type, specialization,  total tokens used)
-- compute node (id, type, provider, gpu_available, price)
-- session (id, userid(FK), agentid (FK), computenodeid(FK), tokens used, retry count, timeElapsed)
--  old table:
-- CREATE TABLE products (
--     id text PRIMARY KEY NOT NULL,
--     name text NOT NULL,
--     description text,
--     price float8)
-- Users (both admins and regular users)
CREATE TYPE user_role AS ENUM (
    'admin',
    'user'
);

CREATE TABLE users (
    id uuid PRIMARY KEY,
    username varchar(50) UNIQUE NOT NULL,
    email varchar(100) UNIQUE NOT NULL,
    encrypted_password varchar(255) NOT NULL, -- bcrypt, argon2, etc. usually 60–100 chars
    ROLE user_role NOT NULL DEFAULT 'user',
    is_active boolean NOT NULL DEFAULT TRUE,
    created_at timestamp NOT NULL DEFAULT NOW(),
    last_login_at timestamp
);

-- AI Agents (different LLM models or agent configurations)
CREATE TABLE ai_agents (
    id uuid PRIMARY KEY,
    name varchar(100) NOT NULL, -- e.g. "Claude 3.5 Sonnet", "Llama 3.1 70B", "Custom RAG Agent"
    model_type varchar(100) NOT NULL, -- e.g. "gpt-4o", "claude-3-5-sonnet", "llama-3.1-70b"
    provider varchar(50) NOT NULL DEFAULT 'openai', -- openai, anthropic, groq, together, local, etc.
    specialization varchar(150), -- "coding", "legal", "creative", NULL = general
    description text,
    context_window integer, -- max context size in tokens
    total_tokens_used bigint NOT NULL DEFAULT 0,
    is_available boolean NOT NULL DEFAULT TRUE,
    created_at timestamp NOT NULL DEFAULT NOW()
);

-- Compute Nodes (instances/machines that run inference)
CREATE TABLE compute_nodes (
    id uuid PRIMARY KEY,
    name varchar(100) NOT NULL, -- e.g. "A100-40GB-01", "H100-80GB-03"
    provider varchar(50) NOT NULL, -- aws, gcp, azure, runpod, vast.ai, self-hosted, ...
    gpu_model_code varchar(100),
    gpu_count integer NOT NULL DEFAULT 1,
    gpu_memory_gb numeric(6, 2),
    price_per_hour numeric(10, 4) NOT NULL, -- USD per hour
    is_online boolean NOT NULL DEFAULT FALSE,
    created_at timestamp NOT NULL DEFAULT NOW()
);

-- Sessions (combination of user + agent + compute node)
CREATE TYPE session_status AS ENUM (
    'pending',
    'running',
    'completed',
    'failed',
    'cancelled',
    'timeout'
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    agent_id uuid NOT NULL REFERENCES ai_agents (id) ON DELETE CASCADE,
    compute_node_id uuid NOT NULL REFERENCES compute_nodes (id) ON DELETE CASCADE,
    status session_status NOT NULL DEFAULT 'pending',
    tokens_used bigint NOT NULL DEFAULT 0,
    retry_count integer NOT NULL DEFAULT 0,
    time_elapsed_ms bigint NOT NULL DEFAULT 0, -- duration in milliseconds
    started_at timestamp,
    ended_at timestamp,
    created_at timestamp NOT NULL DEFAULT NOW()
);

