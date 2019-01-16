#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t lck;

// Note the return type: void*
void* incrementingThreadFunction(){
    for (int j = 0; j < 1000010; j++) {
        pthread_mutex_lock(&lck);
        i++;
        pthread_mutex_unlock(&lck);
    }
    return NULL;
}

void* decrementingThreadFunction(){
    for (int j = 0; j < 1000000; j++) {
        pthread_mutex_lock(&lck);
        i--;
        pthread_mutex_unlock(&lck);
    }
    return NULL;
}


int main(){
    pthread_t incrementingThread, decrementingThread;
    pthread_mutex_init(&lck, NULL);
    
    pthread_create(&incrementingThread, NULL, incrementingThreadFunction, NULL);
    pthread_create(&decrementingThread, NULL, decrementingThreadFunction, NULL);
    
    pthread_join(incrementingThread, NULL);
    pthread_join(decrementingThread, NULL);
    
    printf("The magic number is: %d\n", i);

    pthread_mutex_destroy(&lck);

    return 0;
}
