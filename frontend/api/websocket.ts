import { useEffect, useRef, useState } from "react";

function useWebSocket(url: string, onMessage: (data: any) => void) {
  const ws = useRef<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState(false);

  useEffect(() => {
    ws.current = new WebSocket(url);

    ws.current.onopen = () => {
      setIsConnected(true);
    };

    ws.current.onclose = () => {
      setIsConnected(false);
    };

    ws.current.onmessage = (e) => {
      onMessage(JSON.parse(e.data));
    };

    return () => {
      if (ws.current) {
        ws.current.close();
      }
    };
  }, [url, onMessage]);

  const sendMessage = (message: any) => {
    if (ws.current) {
      ws.current.send(JSON.stringify(message));
    }
  }

  return { sendMessage, isConnected };
}
