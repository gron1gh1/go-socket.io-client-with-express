const express = require('express');
const app = express();
const server = require('http').createServer(app);
const io = require('socket.io')(server);
const cors = require('cors');
const bodyParser = require('body-parser');


app.use(cors());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
app.use(express.static(__dirname+'/asset'));

server.listen(80,()=>console.log('server on 80!'));

io.on('connection',(socket) =>{
  //  console.log(socket);
    socket.on('data',(data) =>{
        console.log('data',data);
        io.emit('data',data);
    })

});

app.get('/',(req,res)=> {
    res.sendFile(__dirname+'/index.html');
});