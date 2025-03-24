import { Route, Routes } from "react-router-dom";
import Blog from "./pages/blog/Blog"
import Index from "./index"

const AppRouter = () => {
  return (
    <Routes>
      <Route path="/" element={<Index />} />
      <Route path="/blog" element={<Blog />} />
    </Routes>
  );
};

export default AppRouter;
