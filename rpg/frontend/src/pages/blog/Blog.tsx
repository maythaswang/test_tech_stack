import { useState, useEffect } from 'react';
import AddPostForm from './AddPostForm';
import PostList from './PostList';
import API_URL from '../../config/apiConfig';
import Post from './Post';

const Blog = () => {
  const [posts, setPosts] = useState<Post[]>([]);

  // Fetch posts from backend
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await fetch(`${API_URL}/api/get_all_messages`);
        const data = await response.json();
        setPosts(data);
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    };
    fetchPosts();
  }, []);

  // Add a new post
  const handleAddPost = async (body: string) => {
    const newPost = { body };
    try {
      const response = await fetch(`${API_URL}/api/post_message`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newPost),
      });

      if (response.ok) {
        const addedPost = await response.json();
        setPosts((prevPosts) => [...prevPosts, addedPost]);
      } else {
        alert('Failed to add post');
      }
    } catch (error) {
      console.error('Error adding post:', error);
    }
  };

  // Delete a post
  const handleDeletePost = async (id: number) => {
    try {
      const response = await fetch(`${API_URL}/api/delete_message/${id}`, {
        method: 'DELETE',
      });

      if (response.ok) {
        setPosts(posts.filter((post) => post.id !== id));
      } else {
        alert('Failed to delete post');
      }
    } catch (error) {
      console.error('Error deleting post:', error);
    }
  };

  return (
    <div className="max-w-4xl mx-auto p-6">
      <h1 className="text-3xl font-bold mb-6">Blog</h1>

      {/* Add New Post Form */}
      <AddPostForm onAddPost={handleAddPost} />

      {/* List of Posts */}
      <PostList posts={posts} onDeletePost={handleDeletePost} />
    </div>
  );
};

export default Blog;
