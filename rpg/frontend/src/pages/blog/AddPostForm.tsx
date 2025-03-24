import React, { useState } from "react";

interface AddPostFormProps {
  onAddPost: (body: string) => void;
}

const AddPostForm = ({ onAddPost }: AddPostFormProps) => {
  const [body, setBody] = useState<string>("");

  // Submit post
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Ensure the user has input some text in the body
    if (body) {
      onAddPost(body);
      setBody("");
    } else {
      alert("Please fill in the body of the post");
    }
  };

  // Form for adding in posts
  return (
    <div className="bg-slate-600 p-6 rounded-lg shadow-lg mb-6">
      <h2 className="text-xl font-bold mb-4">Add New Post</h2>
      <form onSubmit={handleSubmit}>
        <textarea
          placeholder="Write your post..."
          value={body}
          onChange={(e) => setBody(e.target.value)}
          className="w-full p-2 mb-4 border border-gray-300 rounded"
        />
        <button
          type="submit"
          className="w-full py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
        >
          Add Post
        </button>
      </form>
    </div>
  );
};

export default AddPostForm;
