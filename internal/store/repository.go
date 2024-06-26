package store

import "SPORTALK/internal/model"

// Store is the interface that wraps the methods of the store
type UserRepository interface {
	ExistingUser(userName, email string) error
	Login(user *model.User) error
	Register(user *model.User) error
	GetByUUID(uuid string) (*model.User, error)
}

// Store is the interface that wraps the methods of the store
type PostRepository interface {
	Create(post *model.Post) error
	GetAll() ([]*model.Post, error)
	AddCategoryToPost(postID string, categoryID int) error
	GetCategories(postID string) ([]*model.Category, error)
	GetByCategory(categoryID int) ([]*model.Post, error)
}

// Store is the interface that wraps the methods of the store
type CommentRepository interface {
	Create(c *model.Comment) error
	GetByPostID(postID string) ([]*model.Comment, error)
	GetCommentsWithReactionsByPostID(postID string) ([]*model.Comment, error)
}

// Store is the interface that wraps the methods of the store
type ReactionRepository interface {
	CreatePostReaction(reaction *model.Reaction) error
	DeletePostReaction(userID, postID string) error
	GetUserPostReaction(userID, postID string) (*model.Reaction, error)
	CountPostReactions(postID string) (int, error)
	UpdatePostReaction(userID, postID string, reactionID int) error

	CreateCommentReaction(reaction *model.Reaction) error
	DeleteCommentReaction(userID, commentID string) error
	GetUserCommentReaction(userID, commentID string) (*model.Reaction, error)
	CountCommentReactions(commentID string) (int, error)
}
