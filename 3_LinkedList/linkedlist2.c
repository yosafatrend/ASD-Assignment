#include <stdio.h>
#include <stdlib.h>

struct node
{
    int data;
    struct node* next;
};
    typedef struct node node_t;

struct node *head;

void first(){
    struct node *ptr;
    int item;
    ptr = (struct node *) malloc(sizeof(struct node *));
    if (ptr == NULL)
    {
        printf("\n OVERFLOW");
    }else{
        printf("\nInsert value : ");
        scanf("%d", &item);
        ptr->data = item;
        ptr->next = head;
        head = ptr;
        printf("\n Data berhasil disimpan di Node awal");
    }
    
}

void last(){
    struct node *ptr, *temp;
    int item;
    ptr = (struct node *) malloc(sizeof(struct node *));
    if (ptr == NULL)
    {
        printf("\n Overflow");
    }
    else{
        printf("Masukkan data: ");
        scanf("%d", &item);
        ptr->data = item;
        if(head == NULL){
            ptr->next = NULL;
            head = ptr;
            printf("Berhasil disimpan dalam node");
        }
        else{
            temp = head;
            while (temp->next != NULL)
            {
                temp = temp->next;
            }
            temp->next = ptr;
            ptr->next = NULL;
            printf("Data berhasil disimpan di node akhir");
            
        }
    }   
}

void random(){
    int i, loc, item;
    struct node *ptr, *temp;
    ptr = (struct node *)malloc(sizeof(struct node));
    if (ptr == NULL)
    {
        printf("Overflow!");
    }
    else{
        printf("Masukkan data: ");
        scanf("%d", &item);
        ptr->data = item;
        printf("Insert Location : ");
        scanf("%d", &loc);
        temp = head;
        for (i = 0; i < loc; i++)
        {
            temp = temp->next;
            if (temp==NULL)
            {
                printf("\n Tidak dapat tersimpan!");
                return;
            }
        }
        ptr->next = temp->next;
        temp->next = ptr;
        printf("\n NODE berhasil disimpan!");
    }
    
}

void show(){
    struct node *ptr;
    ptr = head;
    if (ptr == NULL)
    {
        printf("Tidak ada data");
    }else{
        printf("Print data ...");
        while (ptr != NULL)
        {
            printf("%d", ptr->data);
            ptr = ptr->next;
        }
        
    }
    
}

int main(){
    int choice = 0;
    while (choice != 5)
    {
        printf("\n ------- Linked List Menu --------\n\n");
        printf("\n 1. Input value to first node");
        printf("\n 2. Input value to last node");
        printf("\n 3. Input value to random node");
        printf("\n 4. Show data");
        printf("\n 5. Exit");
        printf("\n Insert : ");
        scanf(" %d", &choice);

        switch (choice)
        {
            case 1:
            first();
            break;

            case 2:
            last();
            break;

            case 3:
            random();
            break;

            case 4:
            show();
            break;

            case 5:
            exit(0);
            break;
        
            default:
            printf("Silahkan masukkan pilihanmus");
            break;
        }
    }

}