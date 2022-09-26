import axios from "axios";
import React, { createContext, useEffect, useState } from "react";
import BallotComponent from "./components/BallotComponent";
import { BallotContext } from "./context/BallotContext";
import BallotInterface from "./interfaces/ballotInterface"
export const client = axios.create({
  baseURL: "http://localhost:8080"
});

export default function App() {

// const defaultBallot: BallotInterface = { description: "err: Default ballot", id: "wrong", prompts: [] };
// const [currentBallot, setCurrentBallot]: [BallotInterface, (ballots: BallotInterface) => void] = React.useState(defaultBallot);

const defaultBallot: BallotInterface = {id:"", description:"Hello welcome to the robo dance", prompts:[]};

const [currentBallot, setCurrentBallot]: [BallotInterface, (ballots: BallotInterface) => void] = React.useState(defaultBallot);

// const BallotContext = createContext<any>({currentBallot, setCurrentBallot});

  // error and loading states
  const [loading, setLoading]: [boolean, (loading: boolean) => void] = React.useState<boolean>(true);
  const [error, setError]: [string, (error: string) => void] = React.useState("");

  React.useEffect(() => {
    client
      .get<BallotInterface>("/ballots/current", { timeout: 10000 })
      .then(response => {
        setCurrentBallot(response.data);
        setLoading(false);
      })
      .catch(ex => {
        const error =
          ex.code === "ECONNABORTED"
            ? "A timeout has occurred"
            : ex.response.status === 404
              ? "Resource Not found"
              : "An unexpected error has occurred";
        setError(error);
        setLoading(false);
      });
  }, [])

  useEffect(() => {
    const timer = setInterval(()=>{
      client
      .get<BallotInterface>("/ballots/current", { timeout: 10000 })
      .then(response => {
        setCurrentBallot(response.data);
        setLoading(false);
      })
      .catch(ex => {
        const error =
          ex.code === "ECONNABORTED"
            ? "A timeout has occurred"
            : ex.response.status === 404
              ? "Resource Not found"
              : "An unexpected error has occurred";
      });
    }, 5000);
    return () => clearInterval(timer);
  }, []);

  return (<BallotContext.Provider value={{ currentBallot, setCurrentBallot }}>
    <BallotComponent />
    {error && <p className="error">{error}</p>}
    {loading && <p>Loading</p>}
  </BallotContext.Provider>);
  // TODO: add error class
}



