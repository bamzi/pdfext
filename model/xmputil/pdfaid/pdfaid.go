package pdfaid

import (
	_c "fmt"

	_ga "github.com/bamzi/pdfext/model/xmputil/pdfaextension"
	_g "github.com/trimmer-io/go-xmp/xmp"
)

var Namespace = _g.NewNamespace("\u0070\u0064\u0066\u0061\u0069\u0064", "\u0068\u0074\u0074p\u003a\u002f\u002f\u0077w\u0077\u002e\u0061\u0069\u0069\u006d\u002eo\u0072\u0067\u002f\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0069\u0064\u002f", NewModel)

func init() { _g.Register(Namespace, _g.XmpMetadata); _ga.RegisterSchema(Namespace, &Schema) }

// Namespaces implements xmp.Model interface.
func (_da *Model) Namespaces() _g.NamespaceList { return _g.NamespaceList{Namespace} }

// MakeModel gets or create sa new model for PDF/A ID namespace.
func MakeModel(d *_g.Document) (*Model, error) {
	_b, _f := d.MakeModel(Namespace)
	if _f != nil {
		return nil, _f
	}
	return _b.(*Model), nil
}

// SetTag implements xmp.Model interface.
func (_dfg *Model) SetTag(tag, value string) error {
	if _ef := _g.SetNativeField(_dfg, tag, value); _ef != nil {
		return _c.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _ef)
	}
	return nil
}

var _ _g.Model = (*Model)(nil)

// SyncModel implements xmp.Model interface.
func (_ee *Model) SyncModel(d *_g.Document) error { return nil }

// NewModel creates a new model.
func NewModel(name string) _g.Model { return &Model{} }

// SyncToXMP implements xmp.Model interface.
func (_fg *Model) SyncToXMP(d *_g.Document) error { return nil }

// GetTag implements xmp.Model interface.
func (_fd *Model) GetTag(tag string) (string, error) {
	_cg, _ac := _g.GetNativeField(_fd, tag)
	if _ac != nil {
		return "", _c.Errorf("\u0025\u0073\u003a\u0020\u0025\u0076", Namespace.GetName(), _ac)
	}
	return _cg, nil
}

var Schema = _ga.Schema{NamespaceURI: Namespace.URI, Prefix: Namespace.Name, Schema: "\u0050D\u0046/\u0041\u0020\u0049\u0044\u0020\u0053\u0063\u0068\u0065\u006d\u0061", Property: []_ga.Property{{Category: _ga.PropertyCategoryInternal, Description: "\u0050\u0061\u0072\u0074 o\u0066\u0020\u0050\u0044\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "\u0070\u0061\u0072\u0074", ValueType: _ga.ValueTypeNameInteger}, {Category: _ga.PropertyCategoryInternal, Description: "A\u006d\u0065\u006e\u0064\u006d\u0065n\u0074\u0020\u006f\u0066\u0020\u0050\u0044\u0046\u002fA\u0020\u0073\u0074a\u006ed\u0061\u0072\u0064", Name: "\u0061\u006d\u0064", ValueType: _ga.ValueTypeNameText}, {Category: _ga.PropertyCategoryInternal, Description: "C\u006f\u006e\u0066\u006f\u0072\u006da\u006e\u0063\u0065\u0020\u006c\u0065v\u0065\u006c\u0020\u006f\u0066\u0020\u0050D\u0046\u002f\u0041\u0020\u0073\u0074\u0061\u006e\u0064\u0061r\u0064", Name: "c\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0063\u0065", ValueType: _ga.ValueTypeNameText}}, ValueType: nil}

// Can implements xmp.Model interface.
func (_df *Model) Can(nsName string) bool { return Namespace.GetName() == nsName }

// Model is the XMP model for the PdfA metadata.
type Model struct {
	Part        int    `xmp:"pdfaid:part"`
	Conformance string `xmp:"pdfaid:conformance"`
}

// SyncFromXMP implements xmp.Model interface.
func (_bg *Model) SyncFromXMP(d *_g.Document) error { return nil }

// CanTag implements xmp.Model interface.
func (_a *Model) CanTag(tag string) bool { _, _eb := _g.GetNativeField(_a, tag); return _eb == nil }
