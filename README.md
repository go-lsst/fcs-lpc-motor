fcs-lpc-motor
=============

`fcs-lpc-motor` is a simple test program to read (some) values from the Unidrive
M702.

## Installation

```sh
sh> go get github.com/go-lsst/fcs-lpc-motor
```

## Example

```sh
sh> fcs-lpc-motor
[fcs-motor]  === motor [134.158.125.223:502] ===
[fcs-motor] Pr-00.001 (00000): [0 0] 0
[fcs-motor] Pr-00.002 (00001): [117 48] 30000
[fcs-motor] Pr-00.003 (00002): [7 208] 2000
[fcs-motor] Pr-00.004 (00003): [7 208] 2000
[fcs-motor] Pr-00.010 (00009): [0 0] 0
[fcs-motor] Pr-00.031 (00030): [0 0] 0
[fcs-motor] Pr-00.037 (00036): [125 223] 32223
[fcs-motor] Pr-00.048 (00047): [0 3] 3
[fcs-motor] Pr-00.050 (00049): [24 44] 6188
[fcs-motor] Pr-00.016 (00015): [0 150] 150
[fcs-motor] Pr-00.019 (00018): [3 232] 1000
[fcs-motor] Pr-00.020 (00019): [0 1] 1
[fcs-motor] Pr-00.021 (00020): [0 1] 1
[fcs-motor] Pr-00.022 (00021): [0 0] 0
[fcs-motor] Pr-00.023 (00022): [0 0] 0
[fcs-motor] Pr-00.024 (00023): [5 220] 1500
[fcs-motor] Pr-00.025 (00024): [0 200] 200
[fcs-motor] Pr-00.027 (00026): [0 0] 0
[fcs-motor] Pr-00.028 (00027): [0 10] 10
[fcs-motor] Pr-01.006 (00105): [117 48] 30000
[fcs-motor] Pr-01.009 (00108): [0 0] 0
[fcs-motor] Pr-01.010 (00109): [0 1] 1
[fcs-motor] Pr-01.011 (00110): [0 0] 0
[fcs-motor] Pr-01.021 (00120): [0 0] 0
[fcs-motor] Pr-05.001 (00500): [0 0] 0
[fcs-motor] Pr-05.004 (00503): err=modbus: exception '2' (illegal data address), function '131'
[fcs-motor] Pr-05.014 (00513): [0 0] 0
[fcs-motor] Pr-05.053 (00552): [0 0] 0
[fcs-motor] Pr-11.029 (01128): [24 44] 6188
[fcs-motor] Pr-21.015 (02114): [0 0] 0
[fcs-motor] Pr-22.010 (02209): [11 186] 3002
[fcs-motor] Pr-22.011 (02210): [78 62] 20030
[fcs-motor] Pr-22.050 (02249): [43 21] 11029
[fcs-motor] 
[fcs-motor] -- end --
```

Or, interactively:

```sh
sh> fcs-lpc-motor shell
mbus> get 2409
[fcs-motor] Pr-24.010: [0x7d 0xdf] [125 223] (32223)
```
