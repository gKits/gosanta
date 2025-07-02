package mail

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
)

type Body struct {
}

func Send(ctx context.Context) error {
	auth := smtp.PlainAuth("", "", "", "")

	body := new(bytes.Buffer)
	if err := mailTempl().Render(ctx, body); err != nil {
		return fmt.Errorf("[mail] failed to render mail body: %w", err)
	}

	if err := smtp.SendMail("", auth, "", nil, nil); err != nil {
		return fmt.Errorf("[mail] failed to send mail: %w", err)
	}

	return nil
}
