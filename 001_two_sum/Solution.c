#include <stdio.h>
#include <string.h>
#include <malloc.h>

#define MAX_SIZE 2048
int *twoSum(int *nums, int numsSize, int target, int *returnSize)
{
    int hash[MAX_SIZE], *result = malloc(sizeof(int) * 2);
    memset(hash, -1, sizeof(hash));
    for (int i = 0; i < numsSize; i++)
    {
        int tmpIdx = (target - nums[i] + MAX_SIZE) % MAX_SIZE;
        if (hash[tmpIdx] != -1)
        {
            result[0] = hash[tmpIdx];
            result[1] = i;
            *returnSize = 2;
            return result;
        }
        hash[(nums[i] + MAX_SIZE) % MAX_SIZE] = i;
    }

    *returnSize = 0;
    return result;
}

int main()
{
    int nums[] = {2, 7, 11, 15};
    int target = 9;
    int returnSize = 0;
    int *result = twoSum(nums, 4, target, &returnSize);
    printf("%d, %d\n", result[0], result[1]);

    return 0;
}
