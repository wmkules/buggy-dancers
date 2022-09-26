import React, { useEffect, useState } from "react";

export default function App() {
  const [answer, setAnswer] = useState();

  const getAnswer = async () => {
    const res = await fetch("http://139.144.18.143:8080/ballots");
    const data = await res.json();
    setAnswer(data);
  };

  useEffect(() => {
    const timer = setInterval(getAnswer, 2000);
    return () => clearInterval(timer);
  }, []);

  return <div>{JSON.stringify(answer)}</div>;
}



