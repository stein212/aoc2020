// sanity check
const fs = require('fs');

const groups = fs
    .readFileSync('1.in', { encoding: 'utf-8' })
    .split('\n\n')
    .filter((x) => x);

let total = 0;
let part2 = 0;

for (const group of groups) {
    const uniques = new Set([...group.replace(/\n/g, '')]);
    total += uniques.size;

    part2 += [...uniques].filter((char) =>
        group
            .split('\n')
            .filter((x) => x)
            .every((form) => form.includes(char))
    ).length;
}

console.log(total);

console.log(part2);

// let count = 0;
// const qs = [];
// for (let i = 0; i < 26; ++i) qs[i] = 0;

// for (const line of lines) {
//     if (line == '') {
//         for (let i = 0; i < 26; ++i) {
//             if (qs[i] == 1) {
//                 qs[i] = 0;
//                 ++count;
//             }
//         }
//     } else {
//         for (const c of line) {
//             let x = c.charCodeAt(0) - 'a'.charCodeAt(0);
//             if (qs[x] == 0) {
//                 qs[x] = 1;
//             }
//         }
//     }
// }

// for (let i = 0; i < 26; ++i) {
//     if (qs[i] == 1) {
//         qs[i] = 0;
//         ++count;
//     }
// }

// console.log(count);
