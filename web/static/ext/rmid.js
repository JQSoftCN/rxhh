//reset all cells
function doReset(){

}

function doCalc(url){
    var ws = new WebSocket(url);
    ws.onMessage = function(e) {
        var cellVal=eval("("+e.data+")");
        cellId=cellVal.Id;
        var $Cell=$("#"+cellId);
        $Cell.html(cellVal.Val)
    };
}

