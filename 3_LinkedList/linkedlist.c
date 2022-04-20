#include <stdio.h>
#include <stdlib.h>

struct  node
{
    int data;
    struct node* next;
};
typedef struct node node_t;

void printList(node_t *head){
    node_t *temp = head;

    while (temp != NULL){
        printf("%d - ", temp->data);
        temp = temp->next;
    }
    
    printf("\n");
}

int main(){
    node_t n1,n2,n3;
    node_t *head;

    n1.data = 25;
    n2.data = 30;
    n1.data = 20;

    head = &n3;
    n3.next = &n2;
    n2.next = &n1;
    n1.next = NULL;

    printList(head);
}