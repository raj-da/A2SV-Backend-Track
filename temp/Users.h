#include <iostream>
using namespace std;

class UserQueue {
    private:
        class UserNode {
            public:
                string name;
                string bookTitle;
                time_t dueDate;
                UserNode *next;

                UserNode(string name, string bookTitle, time_t dueDate):
                    name(name), bookTitle(bookTitle), dueDate(dueDate), next(nullptr) {}
        };

        UserNode *head = nullptr;
        UserNode *tail = nullptr;

    public:

        // Add a user to the queue
        void enqueue(string name, string bookTitle, time_t dueDate){
            UserNode *newUser = new UserNode(name, bookTitle, dueDate);
            if(!tail)
                head = tail = newUser;
            else{
                tail->next = newUser;
                tail = newUser;
            }
            cout << "User \"" << name << "\" added to the queue for book \"" << bookTitle << "\".\n";
        }

        // Remove a user from the queue
        bool dequeue(string title) {
            if(!head){
                cout << "The queue is empty." << endl;
                return false;
            }

            UserNode *current = head;
            UserNode *previous = nullptr;

            if(current->bookTitle == title){
                head = head->next;
                cout << "User \"" << current->name << "\" removed from the queue for book \"" << current->bookTitle << "\".\n";
                delete current;
                return true;
            }else{
                while(current){
                    if(current->bookTitle == title){
                        previous->next = current->next;
                        cout << "User \"" << current->name << "\" removed from the queue for book \"" << current->bookTitle << "\".\n";
                        delete current;
                        return true;
                    }
                    previous = current;
                    current = current->next;
                }
                return false;
            }

        }

        // get next user in the queue
        UserNode* peek(){
            return head;
        }

        // display the queue
        void displayQueue() {
            if(!head) {
                cout << "The queue is empty." << endl;
                return;
            }

            UserNode *current = head;
            cout << "Queue:\n";
            while(current){
                cout << "Name: " << current->name << ", Book: " << current->bookTitle
                << ", Due Date: " << ctime(&current->dueDate);
                current = current->next;
            }
        }

};