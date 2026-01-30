/*
Section: C
Group Members:
1. Kidist Alemayehu     0771/15
2. Mikiyas Mechalo      0944/15
3. Paul Bereket         1113/15
4. Philipos Hailu       1118/15
5. Rajaf Dereje         1123/15
6. Rebira Oli           1128/15
*/

#include <bits/stdc++.h>
#include "Books_DS.h"
#include <thread>
#include <chrono>
using namespace std;


int main() {
    MyList library;
    int choice;
    bool stop = false;

    do {
        cout << "\nLibrary Management System Menu:\n";
        cout << "1. Add a book to the library\n";
        cout << "2. Display all books in the library\n";
        cout << "3. Search for a book by title\n";
        cout << "4. Search for books by author\n";
        cout << "5. Search for books by year of publication\n";
        cout << "6. Update a book\n";
        cout << "7. Borrow a book\n";
        cout << "8. Return a borrowed book\n";
        cout << "9. Display all borrowed books\n";
        cout << "10. Delete a book by title\n";
        cout << "11. Delete all books by a specific author\n";
        cout << "12. Sort books by author name\n";
        cout << "13. Sort books by year of publication\n";
        cout << "14. Display the total number of books in the library\n";
        cout << "15. Exit the system\n";
        cout << "Enter your choice: ";
        cin >> choice;
        system("cls");

        int t = library.findLength();

        switch (choice) {
            case 1: {
                string title, author;
                int edition, year;
                cin.ignore();
                cout << "Enter the title of the book: ";
                getline(cin, title);
                cout << "Enter the author of the book: ";
                getline(cin, author);
                cout << "Enter the edition of the book: ";
                cin >> edition;
                cout << "Enter the year of publication: ";
                cin >> year;
                library.insertBook(title, author, edition, year);
                cout << "Book added successfully!\n";
                break;
            }
            case 2:
                cout << "Displaying all books in the library:\n";
                library.displayBooks();
                this_thread::sleep_for(chrono::seconds(t));
                break;
            case 3: {
                string title;
                cin.ignore();
                cout << "Enter the title of the book to search: ";
                getline(cin, title);
                library.findBookBytitle(title);
                break;
            }
            case 4: {
                string author;
                cin.ignore();
                cout << "Enter the author to search books: ";
                getline(cin, author);
                library.findBookByauthor(author);
                break;
            }
            case 5: {
                int year;
                cout << "Enter the year to search books: ";
                cin >> year;
                library.findBookByyear(year);
                break;
            }
            case 6: {
                string prevTitle, title, author;
                int edition, year;
                cin.ignore();
                cout << "Enter the title of the book to be updated: ";
                getline(cin, prevTitle);
                cout << "Enter the updated title of the book: ";
                getline(cin, title);
                cout << "Enter the updated author of the book: ";
                getline(cin, author);
                cout << "Enter the updated edition of the book: ";
                cin >> edition;
                cout << "Enter the updated year of publication: ";
                cin >> year;
                
                library.updateBook(prevTitle, title, author, edition, year);
                cout << "Book updated successfully!\n";
                break;
            }
            case 7: {
                string title, userName;
                cin.ignore();
                cout << "Enter your name: ";
                getline(cin, userName);
                cout << "Enter the title of the book you want to borrow: ";
                getline(cin, title);
                library.borrowBook(title, userName);
                this_thread::sleep_for(chrono::seconds(2));
                break;
            }
            case 8: {
                string title;
                cin.ignore();
                cout << "Enter the title of the book to return: ";
                getline(cin, title);
                library.returnBook(title);
                this_thread::sleep_for(chrono::seconds(2));
                break;
            }
            case 9:
                cout << "Displaying all borrowed books:\n";
                library.displayBorrowedBooks();
                this_thread::sleep_for(chrono::seconds(t));
                break;
            case 10: {
                string title;
                cin.ignore();
                cout << "Enter the title of the book to delete: ";
                getline(cin, title);
                library.deleteByTitle(title);
                break;
            }
            case 11: {
                string author;
                cin.ignore();
                cout << "Enter the author to delete all their books: ";
                getline(cin, author);
                library.deleteByAuthor(author);
                break;
            }
            case 12:
                cout << "Sorting books by author name...\n";
                library.sortByAuthor();
                break;
            case 13:
                cout << "Sorting books by year of publication...\n";
                library.sortByYear();
                break;
            case 14:
                cout << "The total number of books in the library: " << library.findLength() << "\n";
                break;
            case 15:
                stop = true;
                cout << "Exiting the system. Thank you!\n";
                break;
            default:
                cout << "Invalid choice! Please select a valid option from the menu.\n";
        }
        this_thread::sleep_for(chrono::seconds(3));
        system("cls");
    } while (!stop);

    return 0;
}
