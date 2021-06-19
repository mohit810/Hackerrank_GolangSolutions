'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', inputStdin => {
    inputString += inputStdin;
});

process.stdin.on('end', _ => {
    inputString = inputString.trim().split('\n').map(string => {
        return string.trim();
    });

    main();
});

function readLine() {
    return inputString[currentLine++];
}

const modifyArray = (nums) => {
    const func = nums.map(function(num) {

        if (num % 2 == 0)
        {
            return 2*num;
        }
        else
        {
            return 3*num;
        }
    });
    return func;
}

