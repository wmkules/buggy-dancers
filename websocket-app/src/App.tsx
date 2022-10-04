import axios from "axios";
import React from "react";
import _ from "lodash";
import BallotComponent from "./components/BallotComponent";
import { BallotContext } from "./context/BallotContext";
import BallotInterface from "./interfaces/ballotInterface";
import Router from "./components/Router";
import "./App.scss";

export default function App() {
  return(
    <div className="container py-4">
    <link rel="preload" href="/fonts/archia-semibold-webfont.woff2" as="font" type="font/woff2" crossOrigin="" />
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap-grid.min.css" />
    <link rel="preload" href="/fonts/archia-medium-webfont.woff2" as="font" type="font/woff2" crossOrigin="" />
    <link rel="preload" href="/fonts/archia-bold-webfont.woff2" as="font" type="font/woff2" crossOrigin="" />    
    <Router />
    </div>
    );
  }