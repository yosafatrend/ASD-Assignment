    #include <stdio.h>
    #include <string.h>
    #define MAX_STACK 5

    typedef struct
    {
        int top;
        char data[5][5];
    }STACK;

    STACK tumpuk;

    void init(){
        tumpuk.top = -1;
    }

    int isFull(){
        if (tumpuk.top == MAX_STACK)
        {
            return 1;
        }else{
            return 0;
        }
        
    }

    int isEmpty(){
        if (tumpuk.top == -1)
        {
            return 1;
        }else{
            return 0;
        }
        
    }

    void Push(char d[5]){
        tumpuk.top++;
        strcpy(tumpuk.data[tumpuk.top], d);
    }

    void Pop(){
        printf("Data yang diambil = %s\n", tumpuk.data[tumpuk.top]);
        tumpuk.top--;
    }

    void tampilStack(){
        for (int i = tumpuk.top; i >= 0; i--)
        {
            printf("Data: %s\n", tumpuk.data[i]);
        }
        
    }

    void Clear(){
        tumpuk.top = -1;
    }

    int main(){
        init();

        int pil;
        char dt[5];

        do
        {
            printf("1. Push\n");
            printf("2. Pop\n");
            printf("3. Cetak isi array stack\n");
            printf("4. Hapus isi array stack\n");
            printf("5. Keluar\n");
            printf("Pilihan anda? ");
            scanf("%d", &pil);

            switch (pil)
            {
                case 1: if (isFull() != 1)
                {
                    printf("Data yang ingin diinputkan = ");
                    scanf("%s", dt);
                    Push(dt);
                }else{
                    printf("Stack sudah penuh!\n");
                }
                    break;
                case 2: if (isEmpty() != 1)
                {
                    Pop();
                }else{
                    printf("Stack masih kosong!\n");
                }
                    break;
                
                case 3: if (isEmpty() != 1)
                {
                    tampilStack();
                }else{
                    printf("\n masih kosong \n");
                }
                break;

                case 4: Clear();
                printf("Data sudah dihapus!\n");
                break;

                default:
                    break;
                }
        } while (pil != 5);
        printf("Keluar dari program\n");
        
    }