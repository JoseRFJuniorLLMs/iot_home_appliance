#ifndef __EDISON_H__
#define __EDISON_H__

#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

// LCD
void initLCD(void);
void setColorLCD(int r, int g, int b);
void writeLCD(int r, int c, char *v);
void removeLCD(void);

#ifdef __cplusplus
}
#endif
#endif
