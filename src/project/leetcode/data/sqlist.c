
#include <stdio.h>
#include "common.h"

#define MAXSIZE 20

typedef struct {
	ElemType *data[MAXSIZE];
	int length;
} SqList;


Status GetElem(SqList *l, int i, ElemType *e) {
	if (l->length == 0 || i < 0 || i > l->length-1) {
		return ERROR;
	}
	*e = *l->data[i-1];
	return OK;
}

Status ListInsert(SqList *l, int i, ElemType *e) {
	if (l->length >= MAXSIZE) {
		return ERROR;
	}
	if (i < 0 || i > l->length) {
		return ERROR;
	}
	if (i < l->length) {
		int k = 0;
		for (k = l->length; k > i; k--) {
			l->data[k] = l->data[k-1];
		}
	}
	l->data[i] = e;
	l->length++;
	return OK;
}

Status ListDelete(SqList *l, int i, ElemType *e) {
	if (l->length == 0) {
		return ERROR;
	}
	if (i < 0 || i > l->length-1) {
		return ERROR;
	}
	*e = *l->data[i];
	int k = 0;
	for (k = i; k < l->length-1; k++) {
		l->data[k] = l->data[k+1];
	}
	l->length--;
	return OK;
}

int main(int argc, char **argv) {
	SqList list = {length: 0};
	ElemType elem = 10;
	ListInsert(&list, 0, &elem);
	ListInsert(&list, 0, &elem);
	ListInsert(&list, 1, &elem);
	ListInsert(&list, 1, &elem);

	printf("length: %d\n", list.length);
	int i = 0;
	for (i = 0; i < list.length; i++) {
		printf("data: %d %d\n", i, *list.data[i]);
	}
}
