create table if not exists user_statuses (
    id          uuid primary key default gen_random_uuid() not null,
    name        varchar(16) not null,
    description text,
);

create table if not exists update_types(
    id         uuid primary key default gen_random_uuid() not null,
    name        varchar(16) not null,
    description text,
)

create table if not exists users (
    id              uuid primary key default gen_random_uuid() not null,
    status_id       uuid references user_statuses(id) not null,
    username        text unique not null,
    password_hash   text not null,
    email           text unique not null,
    otp_secret      text,
    created_at      timestamptz default now() not null,
    updated_at      timestamptz default now() not null,
    created_by      uuid references users(id) not null,
    updated_by     uuid references users(id) not null,
    update_type_id uuid references update_types(id) not null,
);

create table if not exists budget_types (
    id              uuid primary key default gen_random_uuid() not null,
    name            varchar(16) not null,
    description     text,
);

create table if not exists periods(
    id             uuid primary key default gen_random_uuid() not null,
    user_id        uuid references users(id) not null,
    start_date     timestamptz not null,
    end_date       timestamptz not null,
);

create table if not exists entry_types (
    id             uuid primary key default gen_random_uuid() not null,
    name           varchar(16) not null,
    description    text,
);

create table if not exists entries (
    id             uuid primary key default gen_random_uuid() not null,
    budget_id      uuid references budget_types(id) not null,
    entry_type_id  uuid references entry_types(id) not null,
    period_id      uuid references periods(id) not null,
    description    text not null,
    date           timestamptz not null,
    amount         decimal(12,2) not null,
);

create table if not exists budgets(
    id             uuid primary key default gen_random_uuid() not null,
    budget_type_id uuid references budget_types(id) not null,
    user_id        uuid references users(id) not null,
    name           text not null,
    description    text,
    color         varchar(8) default('FFFFFFFF') not null,
    created_at     timestamptz default now() not null,
    updated_at     timestamptz default now() not null,
    created_by     uuid references users(id) not null,
    updated_by     uuid references users(id) not null,
    update_type_id uuid references update_types(id) not null,
);