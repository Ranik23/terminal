package client

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"ssh_connections_manager/internal/entity/errors"
	"ssh_connections_manager/internal/utils/colors"
	"strings"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn        *websocket.Conn
	CurrentDir  string
}

func NewClient(connection *websocket.Conn, dir string) *Client {
	return &Client{
		conn:       connection,
		CurrentDir: dir,
	}
}

func (c *Client) Send(message string) error {
	if err := c.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		return err
	}
	return nil
}

func (c *Client) Execute(ctx context.Context, command string) ([]byte, error) {
	if strings.HasPrefix(command, "cd") {
		parts := strings.Fields(command)

		if len(parts) == 1{
			c.CurrentDir = os.Getenv("HOME")
		} else {
			if err := c.changeDir(parts[1]); err != nil {
				return nil, err
			}
		}
		return []byte(""), nil
	}

	parts := strings.Fields(command)
	args := parts[1:]

	cmd := exec.CommandContext(ctx, parts[0], args...)
	cmd.Dir = c.CurrentDir

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	outputStr := strings.TrimSpace(string(output))

	outputStr = strings.ReplaceAll(outputStr, "\n", " ")

	array := strings.Split(outputStr, " ")

	for i := 0; i < len(array); i++ {
		if c.IsDir(filepath.Join(c.CurrentDir, array[i])) {
			array[i] = colors.ColorString(array[i], colors.ColorCyanBold)
		}
	}

	outputStr = strings.Join(array, " ")

	return []byte(outputStr), nil
}


func (c *Client) changeDir(newDir string) error {
	if newDir == ".." {
		newPath := filepath.Dir(c.CurrentDir) 
		if newPath == c.CurrentDir {
			return errors.ErrAlreadyRoot
		}
		absoluteDir, err := filepath.Abs(newPath)
		if err != nil {
			return errors.ErrInvalidDirectory
		}
		c.CurrentDir = absoluteDir
	} else {
		if strings.HasPrefix(newDir, "/") {
			absoluteDir, err := filepath.Abs(newDir);
			if err != nil {
				return errors.ErrInvalidDirectory
			}
			fileInfo, err := os.Stat(absoluteDir)
			if err != nil {
				return errors.ErrFileStat
			}
			if !fileInfo.IsDir() {
				log.Printf("%s is not a directory", absoluteDir)
				return errors.ErrNotDirectory
			}
			c.CurrentDir = absoluteDir
		} else {
			newPath := filepath.Join(c.CurrentDir, newDir)

			absoluteDir, err := filepath.Abs(newPath);
			if err != nil {
				return errors.ErrInvalidDirectory
			}

			fileInfo, err := os.Stat(absoluteDir)
			if err != nil {
				return errors.ErrFileStat
			}
			if !fileInfo.IsDir() {
				log.Printf("%s is not a directory", absoluteDir)
				return errors.ErrNotDirectory
			}

			c.CurrentDir = absoluteDir
		}
	}

	return nil
}

func (c *Client) ShowDir() string {
	return c.CurrentDir
}


func (c *Client) IsDir(name string) bool {
	fileInfo, err := os.Stat(name)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
