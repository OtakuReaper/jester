-- Creating tables
create table if not exists user_statuses (
    id          uuid primary key default gen_random_uuid() not null,
    name        varchar(16) not null,
    description text
);

create table if not exists users (
    id              uuid primary key default gen_random_uuid() not null,
    status_id       uuid references user_statuses(id) not null,
    username        text unique not null,
    password_hash   text not null,
    email           text unique not null,
    otp_secret      text,
    created_at      timestamptz default now() not null,
    updated_at      timestamptz default now() not null,
    created_by      uuid references users(id),
    updated_by     uuid references users(id)
);

create table if not exists budget_types (
    id              uuid primary key default gen_random_uuid() not null,
    name            varchar(64) not null,
    description     text
);

create table if not exists periods(
    id             uuid primary key default gen_random_uuid() not null,
    user_id        uuid references users(id) not null,
    start_date     timestamptz not null,
    end_date       timestamptz not null
);

create table if not exists entry_types (
    id             uuid primary key default gen_random_uuid() not null,
    name           varchar(16) not null,
    description    text
);

create table if not exists budgets(
    id             uuid primary key default gen_random_uuid() not null,
    user_id        uuid references users(id) not null,
    budget_type_id uuid references budget_types(id) not null,
    name           text not null,
    description    text,
    color          varchar(8) default('FFFFFFFF') not null,
    allocation     decimal(12,2) default 0 not null, -- how much is budgeted
    current_amount decimal(12,2) default 0 not null, -- how is currently remaining in the budget
    created_at     timestamptz default now() not null,
    updated_at     timestamptz default now() not null,
    created_by     uuid references users(id),
    updated_by     uuid references users(id)
);

create table if not exists entries (
    id             uuid primary key default gen_random_uuid() not null,
    budget_id      uuid references budget_types(id) not null,
    entry_type_id  uuid references entry_types(id) not null,
    period_id      uuid references periods(id) not null,
    description    text not null,
    date           timestamptz not null,
    amount         decimal(12,2) not null
);

-- Seeding Tables
insert into user_statuses (name, description) values
('active', 'Active user'),
('inactive', 'Inactive user');

insert into users (status_id, username, password_hash, email) values
(
    (select id from user_statuses where name='active'), 
    'admin', 
    '$2a$12$ZZzXfLNviWQ/UFT1h9.OXOecZHIM9XfzN9.zfuRmXfdL/6lWvNMGe', 
    'example@email.com'
);

insert into budget_types (name, description) values 
('normal', 'Normal budget'),
('fixed', 'Fixed budget'),
('savings', 'Savings budget'),
('income', 'Income budget'),
('debt', 'Debt budget');

insert into entry_types (name, description) values
('debit', 'Money Spent'),
('credit', 'Money Received');

insert into budgets (user_id, budget_type_id, name, description, allocation, current_amount) values 
((select id from users where username='admin'), (select id from budget_types where name='income'), 'Pool', 'Main income budget', 271.23, 0),
((select id from users where username='admin'), (select id from budget_types where name='debt'), 'Land Debt', 'For paying off the Land Debt', 500.00, 0),
((select id from users where username='admin'), (select id from budget_types where name='debt'), 'Mom Debt', 'For paying off the Mom Debt', 44.22, 44.22),
((select id from users where username='admin'), (select id from budget_types where name='debt'), 'Nate Debt', 'For paying off the Laptop', 74.22, 74.22),
((select id from users where username='admin'), (select id from budget_types where name='savings'), 'Emergency Fund', 'For shock and sudden expenses', 20.00, 20.00),
((select id from users where username='admin'), (select id from budget_types where name='fixed'), 'HRCU Contribution', 'For other bank savings', 60.00, 60.00),
((select id from users where username='admin'), (select id from budget_types where name='fixed'), 'Gov. Savings Bank Contribution', 'For other bank savings', 60.00, 60.00),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Credit Card', 'For credit card payments', 62.59, 00.01),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Card Fund', 'For car maintenance', 100.00, 100.00),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Cash', 'For cash withdrawals', 40.00, 0),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Medical', 'For medical expenses', 74.69, 26.76),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Dogs', 'For pet care', 25.00, 25.00),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Car Gas', 'For fueling the cars', 170.00, 120.00),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Groceries', 'For buying groceries', 101.93, 0),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Discretionary Food', 'For eating out and snacks', 100.00, 3.13),
((select id from users where username='admin'), (select id from budget_types where name='fixed'), 'Gym', 'For paying gym membership', 90, 0),
((select id from users where username='admin'), (select id from budget_types where name='savings'), 'Studio Budget', 'For fundings the studio', 42.07, 0),
((select id from users where username='admin'), (select id from budget_types where name='savings'), 'Ranch Budget', 'For fundings the ranch', 42.07, 42.07),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Mom Treat', 'For treating momma', 27.00, 0),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Treating Someone', 'For treating someone', 1.00, 0),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Going Out', 'For spending when I go out', 11.00, 11.00),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Clothing', 'For buying clothes', 11.00, 11.00),
((select id from users where username='admin'), (select id from budget_types where name='normal'), 'Miscellaneous', 'For miscellaneous expenses', 3.00, 0);
