#include <jhd1313m1.h>
#include "edison.hpp"

upm::Jhd1313m1 *lcd;

void initLCD(void) {
    // i2c address: 0x62 RGB, 0x3E LCD
    lcd = new upm::Jhd1313m1(0, 0x3E, 0x62);
}

void setColorLCD(int r, int g, int b) {
    lcd->setColor(r,g,b);
}

void writeLCD(int r, int c, char *v) {
    lcd->setCursor(r,c);
    lcd->write(v);
}

void removeLCD(void) {
    delete lcd;
}
