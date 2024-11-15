CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS todo_lists (
    id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS todo_items (
    id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE
);

/* Таблицы связывающие листы с юзерами многие ко многим */
/* Связь многие ко многим юзеры-листы юзеров*/
CREATE TABLE IF NOT EXISTS users_list (
    id SERIAL PRIMARY KEY, 
    user_id INT NOT NULL,
    list_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE
);

/*Связь многие ко многим АЙтемы - листы (в которых айтемы)*/
CREATE TABLE IF NOT EXISTS item_list (
    id SERIAL NOT NULL PRIMARY KEY,
    list_id INT NOT NULL,
    item_id INT NOT NULL,
    FOREIGN KEY (list_id) REFERENCES todo_lists(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES todo_items(id) ON DELETE CASCADE
);
