CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    phone varchar(20) NOT NULL,
    password varchar(200) NOT NULL,
    registeredAt timestamp DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    status varchar(50) NOT NULL,
    expiresAt timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS lists (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL
); 

CREATE TABLE IF NOT EXISTS users_lists (
    id SERIAL PRIMARY KEY,
    user_id integer NOT NULL,
    list_id integer NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS lists_tasks (
    id SERIAL PRIMARY KEY,
    list_id integer NOT NULL,
    task_id integer NOT NULL,
    FOREIGN KEY (list_id) REFERENCES lists(id) ON DELETE CASCADE,
    FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE
);



    