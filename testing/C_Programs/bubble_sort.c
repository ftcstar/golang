#include<stdio.h>

void bubble_sort(int a[], int len) {
    int i, j, temp;
    for(i=0; i<len; i++){
        for(j=0; j<len-i-1; j++){
            if(a[j] > a[j+1]){
                temp = a[j];
                a[j] = a[j+1];
                a[j+1] = temp;
            }
        }
    }
}

int binary_search(int a[], int l, int h, int k){
    if(h>=l){
        int m = (l+h)/2;
        if(a[m] == k)
            return m;
        if (k<a[m])
            binary_search(a, l, m-1, k);
        binary_search(a, m+1, m-1, k);
    }
    
    return -1;
}

void print_array(int a[], int len){
    for(int i=0; i<len; i++)
        printf("%d\t", a[i]);
    printf("\n");
}

int main(){
    int a[] = {-1, 100, 50, 2, 9, 44};
    int len = sizeof(a)/sizeof(int);
    printf("Arry before sorting\n");
    print_array(a, len);
    bubble_sort(a, len);
    printf("Arry After sorting\n");
    print_array(a, len);
    printf("Search 44 index in array\n");
    int index = binary_search(a, 0, len, 44);
    printf("44 found at index:%d\n", index);
    return 0;
}