#ifndef __EDISON_H__
#define __EDISON_H__

#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

// LCD
void initLCD(void);
void setColor(int r, int g, int b);
void setCursor(int r, int c);
void writeLCD(char *v);
void removeLCD(void);

#ifdef __cplusplus
}
#endif
#endif
