CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name varchar(100),
    email varchar(100),
    phone varchar(20),
    password varchar(200),
    registeredAt timestamp DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    name varchar(255),
    status varchar(50),
    expiresAt timestamp
);

CREATE TABLE IF NOT EXISTS lists (
    id SERIAL PRIMARY KEY,
    name varchar(100) 
); 

CREATE TABLE IF NOT EXISTS users_lists (
    id SERIAL PRIMARY KEY,
    user_id integer,
    list_id integer,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (list_id) REFERENCES lists(id)
);

CREATE TABLE IF NOT EXISTS lists_tasks (
    id SERIAL PRIMARY KEY,
    list_id integer,
    task_id integer,
    FOREIGN KEY (list_id) REFERENCES lists(id),
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);



    