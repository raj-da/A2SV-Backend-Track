#include <bits/stdc++.h>
#include "Users.h"
using namespace std;

class MyList {
private:
    class Node {
    public:
        string title;
        string author;
        int edition;
        int year;
        bool status; // borrowed state availability
        Node* prev;
        Node* next;

        Node(string title, string author, int edition, int year)
            : title(title), author(author), edition(edition), year(year),
              status(false), prev(nullptr), next(nullptr) {}
    };

    Node* head;
    Node* tail;
    UserQueue queue;

public:
    MyList() : head(nullptr), tail(nullptr) {}

    // for writing into the file
    void writeToFile(string title, string author, int edition, int year){
        fstream myFile;
        myFile.open("bookDB.txt", ios::out);
        if(myFile.is_open()){
            myFile << title << '\n';
            myFile << author << '\n';
            myFile << edition << '\n';
            myFile << year << '\n';
        }else{
            cout << "File didn't open";
        }
        myFile.close();
    }

    // // for updating the file
    // void updateFile(string title, string author, int edition, int year){
    //     fstream myFile;
    //     myFile.open("bookDB.txt", ios::out);
    //     if(myFile.is_open()){
    //         myFile << title << '\n';
    //         myFile << author << '\n';
    //         myFile << edition << '\n';
    //         myFile << year << '\n';
    //     }
    //     myFile.close();
    // }

    // Forward traversal
    void displayBooks() {
        Node* current = head;
        if (!current) {
            cout << "The library is empty!" << endl;
            return;
        }

        cout << "-------------------------------------------------------------------\n";
        while (current) {
            cout << "Title: " << current->title << endl;
            cout << "Author: " << current->author << endl;
            cout << "Edition: " << current->edition << endl;
            cout << "Year: " << current->year << endl;
            cout << "####################################################" << endl;
            current = current->next;
        }
        cout << "-------------------------------------------------------------------\n";
    }

    // Count elements in the list
    int findLength() {
        int count = 0;
        Node* current = head;
        while (current) {
            count++;
            current = current->next;
        }
        return count;
    }

    // Search by title
    bool findBookBytitle(const string& title) {
        Node* current = head;
        int position = 1;

        while (current) {
            if (current->title == title) {
                cout << "Book '" << title << "' found at position: " << position << endl;
                return true;
            }
            current = current->next;
            position++;
        }

        cout << "Book '" << title << "' not found in the library." << endl;
        return false;
    }

    // Search by author
    bool findBookByauthor(const string& author) {
        Node* current = head;
        int position = 1;

        while (current) {
            if (current->author == author) {
                cout << "Books by '" << author << "' found starting at position: " << position << endl;
                return true;
            }
            current = current->next;
            position++;
        }

        cout << "No books by '" << author << "' found in the library." << endl;
        return false;
    }

    // Search by year
    bool findBookByyear(int year) {
        Node* current = head;
        int position = 1;

        while (current) {
            if (current->year == year) {
                cout << "Book(s) from year " << year << " found starting at position: " << position << endl;
                return true;
            }
            current = current->next;
            position++;
        }

        cout << "No books from year " << year << " found in the library." << endl;
        return false;
    }

    // Insert a book in sorted order by title
    void insertBook(string title, string author, int edition, int year) {
        Node* new_node = new Node(title, author, edition, year);

        writeToFile(title, author, edition, year);
        
        if (!head) {
            head = tail = new_node;
            return;
        }

        Node* current = head;
        while (current && current->title < title) {
            current = current->next;
        }

        if (!current) {
            tail->next = new_node;
            new_node->prev = tail;
            tail = new_node;
        } else if (current == head) {
            new_node->next = head;
            head->prev = new_node;
            head = new_node;
        } else {
            Node* prev = current->prev;
            prev->next = new_node;
            new_node->prev = prev;
            new_node->next = current;
            current->prev = new_node;
        }

    }

    void updateBook(string prevTitle, string title, string author, int edition, int year) {
        Node* current = head;
        int position = 1;

        while (current) {
            if (current->title == prevTitle) {
                current->title = title;
                current->author = author;
                current->edition = edition;
                current->year = year;
                return;
            }
            current = current->next;
            position++;
        }

        cout << "Book '" << title << "' not found in the library." << endl;
    }

    // Delete by position
    void deleteAtPos(int pos) {
        if (!head) {
            cout << "The library is empty!" << endl;
            return;
        }

        if (pos <= 0 || pos > findLength()) {
            cout << "Invalid position!" << endl;
            return;
        }

        Node* current = head;

        for (int i = 1; i < pos; i++) {
            current = current->next;
        }

        if (current == head) {
            head = head->next;
            if (head) head->prev = nullptr;
        } else if (current == tail) {
            tail = tail->prev;
            if (tail) tail->next = nullptr;
        } else {
            Node* prev = current->prev;
            Node* next = current->next;
            prev->next = next;
            next->prev = prev;
        }

        delete current;
        cout << "Deleted book at position " << pos << "." << endl;
    }

    // Delete by title
    void deleteByTitle(const string& title) {
        if (!head) {
            cout << "The library is empty!" << endl;
            return;
        }

        Node* current = head;
        while (current) {
            if (current->title == title) {
                if (current == head) {
                    head = head->next;
                    if (head) head->prev = nullptr;
                } else if (current == tail) {
                    tail = tail->prev;
                    if (tail) tail->next = nullptr;
                } else {
                    current->prev->next = current->next;
                    current->next->prev = current->prev;
                }
                delete current;
                cout << "Book with title '" << title << "' deleted." << endl;
                return;
            }
            current = current->next;
        }

        cout << "Book with title '" << title << "' not found." << endl;
    }

    // Delete by author
    void deleteByAuthor(const string& author) {
        if (!head) {
            cout << "The library is empty!" << endl;
            return;
        }

        bool deleted = false;
        Node* current = head;

        while (current) {
            if (current->author == author) {
                Node* toDelete = current;
                if (current == head) {
                    head = head->next;
                    if(head) 
                        head->prev = nullptr;
                } else if (current == tail) {
                    tail = tail->prev;
                    if(tail) 
                        tail->next = nullptr;
                } else {
                    current->prev->next = current->next;
                    current->next->prev = current->prev;
                }
                current = current->next;
                delete toDelete;
                deleted = true;
            } else {
                current = current->next;
            }
        }

        if (deleted) {
            cout << "All books by author '" << author << "' deleted." << endl;
        } else {
            cout << "No books by author '" << author << "' found." << endl;
        }
    }

    // for swapping
    void changeVal(string &value1, string &value2) {
        string temp = value1;
        value1 = value2;
        value2 = temp;
    }

    void changeVal(int &value1, int &value2) {
        int temp = value1;
        value1 = value2;
        value2 = temp;
    }

    // Sort books by author
    void sortByAuthor() {
        if (!head || !head->next) return;

        for (Node* i = head; i->next; i = i->next) {
            for (Node* j = i->next; j; j = j->next) {
                if (i->author > j->author) {
                    changeVal(i->title, j->title);
                    changeVal(i->author, j->author);
                    changeVal(i->edition, j->edition);
                    changeVal(i->year, j->year);
                }
            }
        }

        cout << "Books sorted by author." << endl;
    }

    // Sort books by year
    void sortByYear() {
        if (!head || !head->next) return;

        for (Node* i = head; i->next; i = i->next) {
            for (Node* j = i->next; j; j = j->next) {
                if (i->year > j->year) {
                    changeVal(i->title, j->title);
                    changeVal(i->author, j->author);
                    changeVal(i->edition, j->edition);
                    changeVal(i->year, j->year);
                }
            }
        }

        cout << "Books sorted by year." << endl;
    }

    void borrowBook(string title, string userName) {
        Node *current = head;
        while(current) {
            if(current->title == title){
                if(!current->status){
                    time_t now = time(0);
                    time_t dueDate = now + 7 * 24 * 60 * 60; // 7 days in seconds
                    cout << "Book \"" << title << "\" borrowed by \"" << userName
                         << "\". Due Date: " << ctime(&dueDate);
                    current->status = true;
                    return;
                }else{
                    cout << "Book \"" << title << "\" is currently borrowed. Adding \"" << userName << "\" to the queue.\n";
                    time_t now = time(0);
                    queue.enqueue(userName, title, now);
                    return;
                }
            }
            
            current = current->next;
        }
        cout << "Book \"" << title << "\" not found in the library.\n";
    }

    void returnBook(string title){
        Node *current = head;
        while(current){
            if(current->title == title){
                if(current->status){
                    current->status = false;
                    if (queue.dequeue(title)) {
                        cout << "Notifying the next user in the queue.\n";
                    }else{
                        cout << "Book " << title << "returned.\n"; 
                    }
                    return;
                }else{
                    cout << "Book \"" << title << "\" was not borrowed.\n";
                    return; 
                }
            }
            current = current->next;
        }
        cout << "Book \"" << title << "\" not found in the library.\n";
    }

    void displayBorrowedBooks(){
        Node* current = head;
        bool booksFound = false;
        if (!current) {
            cout << "The library is empty!" << endl;
            return;
        }

        cout << "-------------------------------------------------------------------\n";
        while (current) {
            if(current->status){
                cout << "Title: " << current->title << endl;
                cout << "Author: " << current->author << endl;
                cout << "Edition: " << current->edition << endl;
                cout << "Year: " << current->year << endl;
                cout << "####################################################" << endl;
                booksFound = true;
            }
            current = current->next;
        }

        if(!booksFound){
            cout << "No borrowed book";
        }
        cout << "-------------------------------------------------------------------\n";
    }
};