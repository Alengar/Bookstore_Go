import React, { useState } from 'react';

function Bookstore() {
    const [books, setBooks] = useState([]);
    const [newBook, setNewBook] = useState({ title: '', author: '' });

    const addBook = () => {
        setBooks([...books, newBook]);
        setNewBook({ title: '', author: '' });
    };

    return (
        <div>
            <h1>Bookstore</h1>
            <div>
                <h2>Add a Book</h2>
                <input
                    type="text"
                    placeholder="Title"
                    name="title"
                    value={newBook.title}
                    onChange={(e) => setNewBook({ ...newBook, title: e.target.value })}
                />
                <input
                    type="text"
                    placeholder="Author"
                    name="author"
                    value={newBook.author}
                    onChange={(e) => setNewBook({ ...newBook, author: e.target.value })}
                />
                <button onClick={addBook}>Add Book</button>
            </div>
            <div>
                <h2>Books</h2>
                <ul>
                    {books.map((book, index) => (
                        <li key={index}>
                            {book.title} by {book.author}
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
}

export default Bookstore;
