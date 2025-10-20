-- 1. Enable UUID extension
create extension if not exists "uuid-ossp";

-- 2. Create DB tables


-- 2.1 Create the user_roles table and insert default roles
drop table if exists user_roles;

create table if not exists user_roles (
    id uuid primary key default uuid_generate_v4(),
    role_name varchar(50) unique not null,
    description varchar(255) null
);

insert into user_roles (role_name, description) values
('admin', 'Administrator user management and data oversight not using the system normally'),
('user', 'Regular user that uses the system for its intended purpose'),
('suspended', 'removed user who can no longer access the system');

-- 2.2 Create the users table

drop table if exists users;

create table if not exists users (
    id uuid primary key default uuid_generate_v4(),
    role_id uuid references user_roles(id) on delete set null,
    username varchar(50) unique not null,
    email varchar(100) unique not null,
    password_hash varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    created_by uuid references users(id) on delete set null,
    updated_by uuid references users(id) on delete set null
);

-- 2.2.1 Create 2 types of users - admin and regular user
insert into users (role_id, username, email, password_hash) values
((select id from user_roles where role_name = 'admin'), 'admin', 'mikegomez122@gmail.com', 'hashed_password_here'),
((select id from user_roles where role_name = 'user'), 'michael', 'gomezmichaelrobert@gmail.com', 'hashed_password_here');
update users set created_by = (select id from users where username = 'admin') where created_by is null;

-- 2.3 Create categories table

drop table if exists categories;

create table if not exists categories (
    id uuid primary key default uuid_generate_v4(),
    owner_id uuid references users(id) on delete set null,
    name varchar(250) not null,
    description varchar(500) null,

    allotted_amount numeric(10,2) default 0.00 not null,
    current_amount numeric(10,2) default 0.00 not null,
    
    average_spent numeric(10,2) default 0.00 not null,
    total_spent numeric(10,2) default 0.00 not null,
    max_spent numeric(10,2) default 0.00 not null,
    min_spent numeric(10,2) default 0.00 not null,
    
    increase_factor numeric(1,2) default 0.00 not null,
    decrease_factor numeric(1,2) default 0.00 not null,
    same_factor numeric(1,2) default 0.00 not null,
    use_factor numeric(1,2) default 0.00 not null,

    pool_deduction_amount numeric(10,2) default 0.00 not null,

    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    created_by uuid references users(id) on delete set null,
    updated_by uuid references users(id) on delete set null
);

insert into categories (owner_id, name, description, allotted_amount, current_amount, pool_deduction_amount, created_by) values
(
    (select id from users where username = 'michael'), 
    'Groceries', 
    'Dedicated for household resources',
    20.00, 5.00, 20.00, (select id from users where username = 'michael')
),
(
    (select id from users where username = 'michael'), 
    'Discretionary Food', 
    'For discretionary spending on food and dining', 
    45.00, 45.00, 45.00, (select id from users where username = 'michael')
),
(
    (select id from users where username = 'michael'), 
    'Gym', 
    'For gym and fitness-related expenses', 
    45.00, 45.00, 45.00, (select id from users where username = 'michael')
),
(
    (select id from users where username = 'michael'), 
    'Car Gas', 
    'For fueling up the car', 
    100.00, 0.00, 100.00, (select id from users where username = 'michael')
);

-- 2.4 Create periods table

drop table if exists periods;

create table if not exists periods (
    id uuid primary key default uuid_generate_v4(),
    owner_id uuid references users(id) on delete set null,
    name varchar(250) not null,
    start_date date not null,
    end_date date null,
    total_income numeric(10,2) not null,
    amount_used numeric(10,2) default 0.00 not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

-- 2.5 Create entries table

drop table if exists entries;

create table if not exists entries (
    id uuid primary key default uuid_generate_v4(),
    period_id uuid references periods(id) on delete cascade,
    category_id uuid not null,
    description varchar(250) not null,
    entry_date date not null,
    amount numeric(10,2) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

---
