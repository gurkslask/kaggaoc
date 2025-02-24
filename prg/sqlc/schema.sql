CREATE TABLE users (
    user_id INT GENERATED ALWAYS AS IDENTITY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    email TEXT NOT NULL,
    seed TEXT NOT NULL,
  	PRIMARY KEY(user_id)
);


CREATE TABLE completed (
		complete_id INT GENERATED ALWAYS AS IDENTITY,
  	user_id INT,
    challenge int NOT NULL,
  	PRIMARY KEY (complete_id),
    CONSTRAINT fk_user
  		FOREIGN KEY(user_id)
  			REFERENCES users(user_id)
);
