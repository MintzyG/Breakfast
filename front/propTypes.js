import PropTypes from "prop-types";

export const PancakePropTypes = {
  pancake: PropTypes.shape({
    noteId: PropTypes.number.isRequired,
    userId: PropTypes.string,
    emoji: PropTypes.string,
    title: PropTypes.string.isRequired,
    content: PropTypes.string.isRequired,
    color: PropTypes.string,
    createdAt: PropTypes.string,
    updatedAt: PropTypes.string
  }).isRequired,
  onNoteClick: PropTypes.func.isRequired
};
