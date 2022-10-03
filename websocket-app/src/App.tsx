import axios from "axios";
import React from "react";
import _ from "lodash";
import BallotComponent from "./components/BallotComponent";
import { BallotContext } from "./context/BallotContext";
import BallotInterface from "./interfaces/ballotInterface";
import Router from "./components/Router";

export default function App() {
  return(
    <>
      <Router />
    </>
  );
}