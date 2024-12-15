package usecase

import (
	"context"
	"dev-mgmt/internal/domain"
)

type DeviceRepository interface {
	RegisterDevice(ctx context.Context, device *domain.Device) error
	GetDeviceByCert(ctx context.Context, cert string) (*domain.Device, error)
}

type DeviceUseCase struct {
	repo DeviceRepository
}

func NewDeviceUseCase(repo DeviceRepository) *DeviceUseCase {
	return &DeviceUseCase{repo: repo}
}

func (uc *DeviceUseCase) RegisterDevice(ctx context.Context, device *domain.Device) error {
	return uc.repo.RegisterDevice(ctx, device)
}

func (uc *DeviceUseCase) AuthenticateDevice(ctx context.Context, cert string) (*domain.DeviceAuthResult, error) {
	device, err := uc.repo.GetDeviceByCert(ctx, cert)
	if err != nil {
		if err == domain.ErrDeviceNotFound {
			return &domain.DeviceAuthResult{IsValid: false, Reason: "Device not found"}, nil
		}
		return nil, err
	}

	if cert != device.Cert {
		return &domain.DeviceAuthResult{IsValid: false, Reason: "Invalid certificate"}, nil
	}

	return &domain.DeviceAuthResult{IsValid: true, Reason: "Valid certificate"}, nil
}
