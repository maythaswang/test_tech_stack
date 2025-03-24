interface PostItemProps {
  id: number;
  body: string;
  createdAt: string;
  onDelete: (id: number) => void;
}

// Individual Posts
const PostItem = ({ id, body, createdAt, onDelete }: PostItemProps) => {
  return (
    <li className="bg-white p-4 mb-4 rounded-lg shadow-md">
      <p className="text-gray-700">{body}</p>
      <p className="text-sm text-gray-500">
        Created at: {new Date(createdAt).toLocaleString()}
      </p>
      <button
        onClick={() => onDelete(id)}
        className="mt-2 py-1 px-4 bg-red-500 text-white rounded hover:bg-red-600"
      >
        Delete
      </button>
    </li>
  );
};

export default PostItem;
