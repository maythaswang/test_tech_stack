import PostItem from "./PostItem";
import Post from "./Post";

interface PostListProps {
  posts: Post[];
  onDeletePost: (id: number) => void;
}

// List of posts
const PostList = ({ posts, onDeletePost }: PostListProps) => {
  const renderPosts = () => {
    if (!posts) {
      return <p>No posts available.</p>;
    } else {
      return (
        <ul>
          {posts.map((post) => (
            <PostItem
              key={post.id}
              id={post.id}
              body={post.body}
              createdAt={post.created_at}
              onDelete={onDeletePost}
            />
          ))}
        </ul>
      );
    }
  };

  return (
    <div className="mb-6">
      <h2 className="text-xl font-bold mb-4">Posts</h2>
      {renderPosts()}
    </div>
  );
};

export default PostList;
