process.stdin.resume();
process.stdin.setEncoding('ascii');

var input_stdin = "";
var input_stdin_array = "";
var input_currentline = 0;

process.stdin.on('data', function (data) {
    input_stdin += data;
});

process.stdin.on('end', function () {
    input_stdin_array = input_stdin.split("\n");
    main();
});
function readLine() {
    return input_stdin_array[input_currentline++];
}
function Node(data){
    this.data=data;
    this.next=null;
}
function Solution(){

    this.removeDuplicates=function(head){
        let prev = head;

        while (prev) {
            let next = prev.next;

            (next && prev.data == next.data)
                ? (prev.next = next.next)
                : (prev = prev.next);
        }

        return head;
    }

    this.insert=function(head,data){