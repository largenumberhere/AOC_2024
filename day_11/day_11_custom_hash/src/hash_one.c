#include <stdint.h>
#define STONE_T int64_t

#include "../include/hash.h"

// hasher for part one
int hash_stone(STONE_T stone) {
    int hash = -1;
    switch (stone){
        case          125: hash=0;  break;
        case           17: hash=1;  break;
        case       253000: hash=2;  break;
        case            1: hash=3;  break;
        case            7: hash=4;  break;
        case          253: hash=5;  break;
        case            0: hash=6;  break;
        case         2024: hash=7;  break;
        case        14168: hash=8;  break;
        case       512072: hash=9;  break;
        case           20: hash=10; break;
        case           24: hash=11; break;
        case     28676032: hash=12; break;
        case          512: hash=13; break;
        case           72: hash=14; break;
        case            2: hash=15; break;
        case            4: hash=16; break;
        case         2867: hash=17; break;
        case         6032: hash=18; break;
        case      1036288: hash=19; break;
        case         4048: hash=20; break;
        case         8096: hash=21; break;
        case           28: hash=22; break;
        case           67: hash=23; break;
        case           60: hash=24; break;
        case           32: hash=25; break;
        case   2097446912: hash=26; break;
        case           40: hash=27; break;
        case           48: hash=28; break;
        case           80: hash=29; break;
        case           96: hash=30; break;
        case            8: hash=31; break;
        case            6: hash=32; break;
        case            3: hash=33; break;
        case        20974: hash=34; break;
        case        46912: hash=35; break;
        case            9: hash=36; break;
        case        16192: hash=37; break;
        case        12144: hash=38; break;
        case         6072: hash=39; break;
        case     42451376: hash=40; break;
        case     94949888: hash=41; break;
        case        18216: hash=42; break;
        case     32772608: hash=43; break;
        case     24579456: hash=44; break;
        case         4245: hash=45; break;
        case         1376: hash=46; break;
        case         9494: hash=47; break;
        case         9888: hash=48; break;
        case     36869184: hash=49; break;
        case         3277: hash=50; break;
        case         2608: hash=51; break;
        case         2457: hash=52; break;
        case         9456: hash=53; break;
        case           42: hash=54; break;
        case           45: hash=55; break;
        case           13: hash=56; break;
        case           76: hash=57; break;
        case           94: hash=58; break;
        case           98: hash=59; break;
        case           88: hash=60; break;
        case         3686: hash=61; break;
        case         9184: hash=62; break;
        case           77: hash=63; break;
        case           26: hash=64; break;
        case           57: hash=65; break;
        case           56: hash=66; break;
        case            5: hash=67; break;
        case           36: hash=68; break;
        case           86: hash=69; break;
        case           91: hash=70; break;
        case           84: hash=71; break;
        case        10120: hash=72; break;
        case     20482880: hash=73; break;
        case         2048: hash=74; break;
        case         2880: hash=75; break;
        case      5910927: hash=76; break;
        

        default:           hash=-1; break;
    }

    if (hash == -1) {
        printf("hashing failed on %lli\n", stone);
        exit(1);
    }

    return hash;
}

STONE_T unhash_stone(int hash) {
    STONE_T stone = -1;


    switch (hash){
        case 0   : stone =           125;   break; 
        case 1   : stone =            17;   break; 
        case 2   : stone =        253000;   break; 
        case 3   : stone =             1;   break; 
        case 4   : stone =             7;   break; 
        case 5   : stone =           253;   break; 
        case 6   : stone =             0;   break; 
        case 7   : stone =          2024;   break; 
        case 8   : stone =         14168;   break; 
        case 9   : stone =        512072;   break; 
        case 10  : stone =            20;   break;
        case 11  : stone =            24;   break;
        case 12  : stone =      28676032;   break;
        case 13  : stone =           512;   break;
        case 14  : stone =            72;   break;
        case 15  : stone =             2;   break;
        case 16  : stone =             4;   break;
        case 17  : stone =          2867;   break;
        case 18  : stone =          6032;   break;
        case 19  : stone =       1036288;   break;
        case 20  : stone =          4048;   break;
        case 21  : stone =          8096;   break;
        case 22  : stone =            28;   break;
        case 23  : stone =            67;   break;
        case 24  : stone =            60;   break;
        case 25  : stone =            32;   break;
        case 26  : stone =    2097446912;   break;
        case 27  : stone =            40;   break;
        case 28  : stone =            48;   break;
        case 29  : stone =            80;   break;
        case 30  : stone =            96;   break;
        case 31  : stone =             8;   break;
        case 32  : stone =             6;   break;
        case 33  : stone =             3;   break;
        case 34  : stone =         20974;   break;
        case 35  : stone =         46912;   break;
        case 36  : stone =             9;   break;
        case 37  : stone =         16192;   break;
        case 38  : stone =         12144;   break;
        case 39  : stone =          6072;   break;
        case 40  : stone =      42451376;   break;
        case 41  : stone =      94949888;   break;
        case 42  : stone =         18216;   break;
        case 43  : stone =      32772608;   break;
        case 44  : stone =      24579456;   break;
        case 45  : stone =          4245;   break;
        case 46  : stone =          1376;   break;
        case 47  : stone =          9494;   break;
        case 48  : stone =          9888;   break;
        case 49  : stone =      36869184;   break;
        case 50  : stone =          3277;   break;
        case 51  : stone =          2608;   break;
        case 52  : stone =          2457;   break;
        case 53  : stone =          9456;   break;
        case 54  : stone =            42;   break;
        case 55  : stone =            45;   break;
        case 56  : stone =            13;   break;
        case 57  : stone =            76;   break;
        case 58  : stone =            94;   break;
        case 59  : stone =            98;   break;
        case 60  : stone =            88;   break;
        case 61  : stone =          3686;   break;
        case 62  : stone =          9184;   break;
        case 63  : stone =            77;   break;
        case 64  : stone =            26;   break;
        case 65  : stone =            57;   break;
        case 66  : stone =            56;   break;
        case 67  : stone =             5;   break;
        case 68  : stone =            36;   break;
        case 69  : stone =            86;   break;
        case 70  : stone =            91;   break;
        case 71  : stone =            84;   break;
        case 72  : stone =         10120;   break;
        case 73  : stone =      20482880;   break;
        case 74  : stone =          2048;   break;
        case 75  : stone =          2880;   break;
        default  : stone =            -1;   break;  
    }
    if (stone == -1) {
        perror("invalid hash");
        exit(1);
    }

    return stone;
}