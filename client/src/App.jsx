import { Routes, Route } from "react-router-dom";
import Home from './page/Home';
import Blog from './page/Blog';
import Header from "./components/layouts/Header";
import Footer from "./components/layouts/Footer";

function App() {

  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/blog/:id" element={<Blog />} />
      </Routes>
      <Footer />
    </>
  );
}

export default App
