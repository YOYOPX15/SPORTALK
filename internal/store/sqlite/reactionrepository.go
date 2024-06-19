package sqlite

import (
	"Forum/internal/model"
	"fmt"
)

type ReactionRepository struct {
	store *Store
}

func (r *ReactionRepository) CreateCommentReaction(reaction *model.Reaction) error {
	queryInsert := "INSERT INTO comment_reactions(comment_id, user_UUID, reaction_id) VALUES (?, ?, ?)"
	_, err := r.store.Db.Exec(queryInsert, reaction.CommentID, reaction.UserID, reaction.ReactionID)
	if err != nil {
		return fmt.Errorf("createCommentReaction error: %w", err)
	}
	return nil
}
func (r *ReactionRepository) UpdatePostReaction(userID, postID string, reactionID int) error {
	queryUpdate := "UPDATE post_reactions SET reaction_id = ? WHERE user_UUID = ? AND post_id = ?"
	_, err := r.store.Db.Exec(queryUpdate, reactionID, userID, postID)
	if err != nil {
		return err
	}
	return nil
}
func (r *ReactionRepository) DeleteCommentReaction(userID, commentID string) error {
	queryDelete := "DELETE FROM comment_reactions WHERE user_UUID = ? AND comment_id = ?"
	_, err := r.store.Db.Exec(queryDelete, userID, commentID)
	return err
}

func (r *ReactionRepository) GetUserCommentReaction(userID, commentID string) (*model.Reaction, error) {
	var reaction model.Reaction
	queryGet := "SELECT user_UUID, comment_id, reaction_id FROM comment_reactions WHERE user_UUID = ? AND comment_id = ?"
	err := r.store.Db.QueryRow(queryGet, userID, commentID).Scan(&reaction.UserID, &reaction.CommentID, &reaction.ReactionID)
	return &reaction, err
}

func (r *ReactionRepository) CountCommentReactions(commentID string) (int, error) {
	queryCount := "SELECT COUNT(*) FROM comment_reactions WHERE comment_id = ?"
	var count int
	err := r.store.Db.QueryRow(queryCount, commentID).Scan(&count)
	return count, err
}

func (r *ReactionRepository) CreatePostReaction(reaction *model.Reaction) error {
	queryInsert := "INSERT INTO post_reactions(post_id, user_UUID, reaction_id) VALUES (?, ?, ?)"
	_, err := r.store.Db.Exec(queryInsert, reaction.PostID, reaction.UserID, reaction.ReactionID)
	if err != nil {
		return fmt.Errorf("createReaction error: %w", err)
	}
	return nil
}

func (r *ReactionRepository) DeletePostReaction(userID, postID string) error {
	queryDelete := "DELETE FROM post_reactions WHERE user_UUID = ? AND post_id = ?"
	_, err := r.store.Db.Exec(queryDelete, userID, postID)
	return err
}

func (r *ReactionRepository) GetUserPostReaction(userID, postID string) (*model.Reaction, error) {
	var reaction model.Reaction
	queryGet := "SELECT user_UUID, post_id, reaction_id FROM post_reactions WHERE user_UUID = ? AND post_id = ?"
	err := r.store.Db.QueryRow(queryGet, userID, postID).Scan(&reaction.UserID, &reaction.PostID, &reaction.ReactionID)
	return &reaction, err
}

func (r *ReactionRepository) CountPostReactions(postID string) (int, error) {
	queryCount := "SELECT COUNT(*) FROM post_reactions WHERE post_id = ?"
	var count int
	err := r.store.Db.QueryRow(queryCount, postID).Scan(&count)
	return count, err
}
