
CREATE TABLE book (
        book_id UUID NOT NULL UNIQUE,
        name VARCHAR NOT NULL,
        author_name VARCHAR,
        price INTEGER NOT NULL,
        date VARCHAR NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL 
);

CREATE TABLE users (
        user_id UUID NOT NULL UNIQUE,
        first_name VARCHAR NOT NULL,
        last_name VARCHAR NOT NULL,
        phone_number VARCHAR NOT NULL,
        balance INTEGER, 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL 
);

CREATE TABLE orders (
        order_id UUID NOT NULL,
        books_id UUID NOT NULL REFERENCES book(book_id),
        users_id UUID NOT NULL REFERENCES users(user_id),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);