using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Net.Sockets;
using System.Net;
using System.Text;
using System;

public class GameData
{
    public int Count { get; set; }
}

public class NetworkManager : MonoBehaviour
{
    Socket clientSocket;
    int count = 0;

    void Start()
    {
        // initialize client socket
        clientSocket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.IP);
        clientSocket.Connect(new IPEndPoint(IPAddress.Loopback, 7000));
    }

    private void Update()
    {

        // send string data
        byte[] data = Encoding.Default.GetBytes("this is unity" + (count++)%10);
        clientSocket.BeginSend(data, 0, 14, SocketFlags.None,
                                new AsyncCallback(sendStr), clientSocket);

    }

    static void sendStr(IAsyncResult ar)
    {
        Socket clientSocket = (Socket)ar.AsyncState;
        //int strLength = clientSocket.EndSend(ar);
    }
}
