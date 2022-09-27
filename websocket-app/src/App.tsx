import axios from "axios";
import React from "react";
import _ from "lodash"
import BallotComponent from "./components/BallotComponent";
import { BallotContext } from "./context/BallotContext";
import BallotInterface from "./interfaces/ballotInterface"
export const client = axios.create({
  baseURL: "http://localhost:8080"
});

export default function App() {

  const defaultBallot: BallotInterface = { id: "", description: "Hello welcome to the robo dance", prompts: [] };

  const [currentBallot, _setCurrentBallot]: [BallotInterface, any] = React.useState(defaultBallot);

  const setCurrentBallot = (bal:BallotInterface) => {
    _setCurrentBallot((prev)=>{
      if (!_.isEqual(bal, prev)){
        return bal;
      }
      return prev;
    })
  }
  // error and loading states
  const [loading, setLoading]: [boolean, (loading: boolean) => void] = React.useState<boolean>(true);
  const [error, setError]: [string, (error: string) => void] = React.useState("");


function getCurrentBallot() {
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

}

  React.useEffect(() => {
    getCurrentBallot()
    const timer = setInterval(() => getCurrentBallot(), 2000);
    return () => clearInterval(timer);
  }, [])

  return (<BallotContext.Provider value={{ currentBallot, setCurrentBallot }}>
    <BallotComponent />
    {error && <p className="error">{error}</p>}
    {loading && <p>Loading</p>}
  </BallotContext.Provider>);
  // TODO: add error class
}
