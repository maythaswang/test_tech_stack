import PostItem from './PostItem';
import Post from './Post'

interface PostListProps {
  posts: Post[];
  onDeletePost: (id: number) => void;
}

// List of posts
const PostList = ({ posts, onDeletePost }: PostListProps) => (
  <div className="mb-6">
    <h2 className="text-xl font-bold mb-4">Posts</h2>

    {/* Check if any posts exist */}
    {posts.length === 0 ? (
      <p>No posts available.</p>
    ) : (
      // Show when posts exist
      <ul>
        {posts.map((post) => (
          <PostItem
            key={post.id}
            id={post.id}
            body={post.body}
            createdAt={post.createdAt}
            onDelete={onDeletePost}
          />
        ))}
      </ul>
    )}
  </div>
);

export default PostList;