package redis

import "context"

type (
	BfInsertArgs struct {
		Key      string
		Values   []interface{}
		Capacity uint32
		//NoCreate bool
	}
)

func (c *Client) BfInsert(ctx context.Context, op *BfInsertArgs) *Cmd {
	args := make([]interface{}, 1, 3+len(op.Values))
	args[0] = op.Key
	cap := op.Capacity
	if cap <= 100 {
		cap = 10000
	}
	args[1] = "CAPACITY"
	args[2] = cap
	args = appendArgs(args, op.Values)
	return c.Do(ctx, "BF.INSERT", args)
}

func (c *Client) BfAdd(ctx context.Context, key string, value interface{}) *Cmd {

	return c.Do(ctx, "BF.ADD", key, value)

}
func (c *Client) BfMadd(ctx context.Context, key string, values ...interface{}) *Cmd {

	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = appendArgs(args, values)
	return c.Do(ctx, "BF.MADD", args)

}

func (c *Client) BfExists(ctx context.Context, key string, value interface{}) *Cmd {

	return c.Do(ctx, "BF.EXISTS", key, value)

}
func (c *Client) BfMexists(ctx context.Context, key string, values ...interface{}) *Cmd {

	args := make([]interface{}, 1, 1+len(values))
	args[0] = key
	args = appendArgs(args, values)
	return c.Do(ctx, "BF.MEXISTS", args)

}
