import ReactDOM from "react-dom";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import VisHelper from "../pages/VisHelper";
import Voting from "../pages/Voting";
import Dashboard from "../pages/Dashboard";

export default function Router() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/voting" element={<Voting />} />
                <Route path="/vishelper" element={<VisHelper />} />
                <Route path="/dashboard" element={<Dashboard />} />
            </Routes>
        </BrowserRouter>
    );
}