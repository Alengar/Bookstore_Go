import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import Bookstore from './Bookstore';

test('renders the bookstore UI', () => {
    const { getByText, getByPlaceholderText } = render(<Bookstore />);
    const titleInput = getByPlaceholderText('Title');
    const authorInput = getByPlaceholderText('Author');
    const addButton = getByText('Add Book');

    expect(titleInput).toBeInTheDocument();
    expect(authorInput).toBeInTheDocument();
    expect(addButton).toBeInTheDocument();
});

test('allows adding a book', () => {
    const { getByPlaceholderText, getByText, getByTestId } = render(<Bookstore />);
    const titleInput = getByPlaceholderText('Title');
    const authorInput = getByPlaceholderText('Author');
    const addButton = getByText('Add Book');
    const booksList = getByTestId('books-list');

    fireEvent.change(titleInput, { target: { value: 'Sample Book' } });
    fireEvent.change(authorInput, { target: { value: 'Author Name' } });
    fireEvent.click(addButton);

    expect(booksList).toHaveTextContent('Sample Book by Author Name');
});
