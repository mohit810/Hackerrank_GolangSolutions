function action(e) {
    let btn = e.target || e.srcElement;
    let actionButton = document.getElementById(btn.id).innerHTML;
    let res = document.getElementById('res');

    switch(actionButton) {
        case '0':
        case '1':
        case '+':
        case '-':
        case '*':
        case '/':
            res.innerHTML += actionButton;
            break;
        case 'C':
            res.innerHTML = '';
            break;
        case '=':
            let expr = res.innerHTML;
            let nums = /(\d+)/g;

            expr = expr.replace(nums, function(match) {
                return parseInt(match, 2);
            })

            res.innerHTML = eval(expr).toString(2);
            break;
        default:
            break;
    }
}

var buttons = document.getElementsByTagName('button');
for (let button of buttons) {
    button.onclick = action;
}