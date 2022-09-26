import axios from "axios";
import React, { useEffect, useState } from "react";
import BallotComponent from "./components/BallotComponent";
import BallotInterface from "./interfaces/ballotInterface"

export default function App() {
  const defaultBallots: BallotInterface[] = [];

  const [allBallots, setAllBallots]: [BallotInterface[], (ballots: BallotInterface[]) => void] = React.useState(defaultBallots);

  React.useEffect(() => {
    axios
      .get<BallotInterface[]>("http://localhost:8080/ballots")
      .then(response => {
        setAllBallots(response.data);
      })
      .catch(ex => {
        const error =
          ex.response.status === 404
            ? "Resource Not found"
            : "An unexpected error has occurred";
      });
  }, [])

  return <div>
    {allBallots.map((b) => <BallotComponent id={b.id} description={b.description} prompts={b.prompts} />)}
  </div>;
}



