package oddb

import (
	"errors"
	"io"
)

// ErrRecordNotFound is returned from Get and Delete when Database
// cannot find the Record by the specified key
var ErrRecordNotFound = errors.New("oddb: Record not found for the specified key")

// Database represents a collection of record (either public or private)
// in a container.
//
// TODO: We might need to define standard errors for common failures
// of database operations like ErrRecordNotFound
type Database interface {
	// ID returns the identifier of the Database.
	ID() string

	// Get fetches the Record identified by the supplied key and
	// writes it onto the supplied Record.
	//
	// Get returns an ErrRecordNotFound if Record identified by
	// the supplied key does not exist in the Database.
	// It also returns error if the underlying implementation
	// failed to read the Record.
	Get(key string, record *Record) error

	// Save updates the supplied Record in the Database if Record with
	// the same key exists, else such Record is created.
	//
	// Save returns an error if the underlying implemention failed to
	// create / modify the Record.
	Save(record *Record) error

	// Delete removes the Record identified by the key in the Database.
	//
	// Delete returns an ErrRecordNotFound if the Record identified by
	// the supplied key does not exist in the Database.
	// It also returns an error if the underlying implementation
	// failed to remove the Record.
	Delete(key string) error

	// Query executes the supplied query against the Database and returns
	// an Rows to iterate the results.
	Query(query *Query) (*Rows, error)

	GetSubscription(key string, subscription *Subscription) error
	SaveSubscription(subscription *Subscription) error
	DeleteSubscription(key string) error
}

// Rows implements a scanner-like interface for easy iteration on a
// result set returned from a query
type Rows struct {
	iter    RowsIter
	lasterr error
	closed  bool
	record  Record
	nexted  bool
}

// NewRows creates a new Rows.
//
// Driver implementators are expected to call this method with
// their implementation of RowsIter to return a Rows from Database.Query.
func NewRows(iter RowsIter) *Rows {
	return &Rows{
		iter: iter,
	}
}

// Close closes the Rows and prevents futher enumerations on the instance.
func (r *Rows) Close() error {
	if r.closed {
		return nil
	}

	r.closed = true
	return r.iter.Close()
}

// Scan tries to prepare the next record and returns whether such record
// is ready to be read.
func (r *Rows) Scan() bool {
	if r.closed {
		return false
	}

	r.lasterr = r.iter.Next(&r.record)
	if r.lasterr != nil {
		r.Close()
		return false
	}

	return true
}

// Record returns the current record in Rows.
//
// It must be called after calling Scan and Scan returned true.
// If Scan is not called or previous Scan return false, the behaviour
// of Record is unspecified.
func (r *Rows) Record() Record {
	return r.record
}

// Err returns the last error encountered during Scan.
//
// NOTE: It is not an error if the underlying result set is exhausted.
func (r *Rows) Err() error {
	if r.lasterr == io.EOF {
		return nil
	}

	return r.lasterr
}

// RowsIter is an iterator on results returned by execution of a query.
type RowsIter interface {
	// Close closes the rows iterator
	Close() error

	// Next populates the next Record in the current rows iterator into
	// the provided record.
	//
	// Next should return io.EOF when there are no more rows
	Next(record *Record) error
}