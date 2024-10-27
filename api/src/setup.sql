DROP DATABASE IF EXISTS breakfast_db;

CREATE DATABASE breakfast_db WITH TEMPLATE=template0 OWNER=postgres;
\connect breakfast_db;

CREATE TABLE users (
  id UUID PRIMARY KEY UNIQUE NOT NULL,
  first_name VARCHAR(31) NOT NULL,
  last_name VARCHAR(31) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password TEXT NOT NULL
);

-- Categories for sorting
CREATE TABLE categories (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  title VARCHAR(127) NOT NULL UNIQUE,
  description VARCHAR(255),
  emoji VARCHAR(15) NOT NULL,
  color CHAR(7) NOT NULL,
  text_color CHAR(7) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Yogurt - Tasks
CREATE TABLE yogurt (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  emoji VARCHAR(15) NOT NULL,
  title VARCHAR(127) NOT NULL,
  description VARCHAR(255),
  completed BOOLEAN NOT NULL,
  difficulty INT NOT NULL,
  task_size INT NOT NULL,
  priority INT NOT NULL,
  category_id INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Toast - Time Tracking
CREATE TABLE toast (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  session_name VARCHAR(127) NOT NULL,
  description VARCHAR(255),
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP,
  duration INTERVAL,
  category_id INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Pancake - Notes/Journaling
CREATE TABLE pancake (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  title VARCHAR(127) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  category_id INTEGER,
  tags TEXT[],
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Cereal - Scheduling
CREATE TABLE cereal (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  title VARCHAR(127) NOT NULL,
  description VARCHAR(255),
  start_time TIMESTAMP NOT NULL,
  end_time TIMESTAMP NOT NULL,
  location VARCHAR(255),
  is_recurring BOOLEAN NOT NULL DEFAULT FALSE,
  recurrence_pattern VARCHAR(50),
  category_id INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Espresso - Focus Timer
CREATE TABLE espresso (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  session_name VARCHAR(127) NOT NULL,
  focus_duration INT NOT NULL, -- Duration in minutes
  break_duration INT NOT NULL, -- Duration in minutes
  completed_cycles INT NOT NULL DEFAULT 0,
  target_cycles INT NOT NULL,
  started_at TIMESTAMP,
  completed_at TIMESTAMP,
  category_id INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Omelette - Kanban Board
CREATE TABLE omelette_boards (
  id SERIAL PRIMARY KEY NOT NULL,
  user_id UUID NOT NULL,
  title VARCHAR(127) NOT NULL,
  description VARCHAR(255),
  category_id INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE omelette_columns (
  id SERIAL PRIMARY KEY NOT NULL,
  board_id INT NOT NULL,
  title VARCHAR(127) NOT NULL,
  position INT NOT NULL,
  FOREIGN KEY (board_id) REFERENCES omelette_boards(id)
);

CREATE TABLE omelette_cards (
  id SERIAL PRIMARY KEY NOT NULL,
  column_id INT NOT NULL,
  title VARCHAR(127) NOT NULL,
  description VARCHAR(255),
  position INT NOT NULL,
  due_date TIMESTAMP,
  FOREIGN KEY (column_id) REFERENCES omelette_columns(id)
);

-- Create indexes for better performance
CREATE INDEX idx_yogurt_user_id ON yogurt(user_id);
CREATE INDEX idx_toast_user_id ON toast(user_id);
CREATE INDEX idx_omelette_boards_user_id ON omelette_boards(user_id);
CREATE INDEX idx_pancake_user_id ON pancake(user_id);
CREATE INDEX idx_cereal_user_id ON cereal(user_id);
CREATE INDEX idx_espresso_user_id ON espresso(user_id);

CREATE OR REPLACE FUNCTION delete_category(
  p_user_id UUID,
  p_category_id INTEGER
) RETURNS BOOLEAN AS $$
DECLARE
  v_category_exists BOOLEAN;
BEGIN
  SELECT EXISTS (
    SELECT 1 FROM categories
    WHERE id = p_category_id AND user_id = p_user_id
  ) INTO v_category_exists;

  IF NOT v_category_exists THEN
    RETURN FALSE;
  END IF;

  UPDATE yogurt
  SET category_id = NULL
  WHERE category_id = p_category_id AND user_id = p_user_id;

  UPDATE toast
  SET category_id = NULL
  WHERE category_id = p_category_id AND user_id = p_user_id;

  UPDATE omelette_boards
  SET category_id = NULL
  WHERE category_id = p_category_id AND user_id = p_user_id;
 
  UPDATE pancake
  SET category_id = NULL
  WHERE category_id = p_category_id AND user_id = p_user_id;
 
  UPDATE cereal
  SET category_id = NULL
  WHERE category_id = p_category_id AND user_id = p_user_id;
 
  UPDATE espresso
  SET category_id = NULL
  WHERE category_id = p_category_id AND user_id = p_user_id;
 
  DELETE FROM categories
  WHERE id = p_category_id AND user_id = p_user_id;
 
  RETURN TRUE;
END;
$$ LANGUAGE plpgsql;
