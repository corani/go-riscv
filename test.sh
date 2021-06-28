#!/bin/bash

failed=0
passed=0

for name in ./riscv-tests/isa/rv32ui-p-*.dump; do
    name=$(basename $name)
    echo "emulate: ${name%.*}"
    ./bin/emulate -gas 530 -in "./riscv-tests/isa/${name%.*}" > /tmp/${name%.*}.out
    code=$?
    if [[ $code == 0 ]]; then
        echo "PASS"

        ((passed++))
    else
        echo "FAIL: ${code}"

        ((failed++))
    fi
done

echo "Totals: pass=${passed} / fail=${failed}"
exit $failed
