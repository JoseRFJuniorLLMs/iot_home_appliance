#include "edison.hpp"

void initLCD(void) {
    // i2c address: 0x62 RGB, 0x3E LCD
    lcd = new upm::Jhd1313m1(0, 0x3E, 0x62);
}

void setColor(int r, int g, int b) {
    lcd->setColor(r,g,b);
}

void setCursor(int r, int c) {
    lcd->setCursor(r,c);
}

void writeLCD(char *v) {
    lcd->write(v);
}

void removeLCD(void) {
    delete lcd;
}
